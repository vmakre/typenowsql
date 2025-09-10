
GO
/****** Object:  StoredProcedure [dbo].[sp_ExecuteStoredQueryWithJSON]    Script Date: 10.9.2025 12:47:37 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
-- Enhanced execution procedure with JSON parameter handling
ALTER PROCEDURE [dbo].[sp_ExecuteStoredQueryWithJSON]
    @QueryName NVARCHAR(100),
    @ParamInput NVARCHAR(MAX) = NULL
--	@MetaInfo BIT = 0
AS
BEGIN
    SET NOCOUNT ON;

    DECLARE @QueryText NVARCHAR(MAX);
    DECLARE @ParamDefinition NVARCHAR(MAX) = '';
    DECLARE @SQL NVARCHAR(MAX);
    DECLARE @ParamList NVARCHAR(MAX) = '';

    
    -- Get the query text and parameter definitions 
    SELECT @QueryText = QueryText, @ParamDefinition = Parameters  
    FROM dbo.StoredQueries
    WHERE QueryName = @QueryName AND IsActive = 1;
    
    IF @QueryText IS NULL
    BEGIN
        RAISERROR('Query not found or inactive', 16, 1);
        RETURN;
    END
    -- Parse parameters if provided
    IF @ParamInput IS NOT NULL AND @ParamDefinition IS NOT NULL
    BEGIN
		SELECT @ParamList = STUFF(
			(
				SELECT ', ' + 
					   '@' + section.[key] + ' '  +
					    CASE 
							WHEN JSON_VALUE(section.value, '$.type') = 'int' THEN 'INT'
							WHEN JSON_VALUE(section.value, '$.type') = 'date' THEN 'DATE'
							WHEN JSON_VALUE(section.value, '$.type') = 'varchar' THEN 'NVARCHAR(MAX)'
							ELSE 'NVARCHAR(MAX)'
					    END
				FROM OPENJSON(@ParamDefinition) as section
				FOR XML PATH(''), TYPE
			).value('.', 'NVARCHAR(MAX)'), 
			1, 2, '')
		-- PRINT  @ParamList 
        -- Build the parameter values assignment
        DECLARE @ParamVBuild NVARCHAR(MAX) = '';
		SELECT @ParamVBuild = STUFF(
				(
			SELECT ', ' + 
				   '@' + [key] + ' = ' + 
					   CASE 
						--   WHEN json_value(@ParamDefinition, '$.' + [key] + '.type') IN ('int', 'decimal') THEN [value]
							WHEN json_value(@ParamDefinition, '$."'+ [key]+'".type' ) IN ('int', 'decimal') THEN [value]
						   ELSE '''' + REPLACE([value], '''', '''''') + ''''
					   END
						FROM OPENJSON(@ParamInput) 
						FOR XML PATH(''), TYPE
					).value('.', 'NVARCHAR(MAX)'), 
					1, 2, '')
	    -- PRINT @ParamVBuild
		SET @SQL = N'EXEC sp_executesql @QueryText, @ParamList, ' + @ParamVBuild;
		BEGIN
			-- Execute with parameters
			SET @SQL = N'EXEC sp_executesql @QueryText, @ParamList, ' + @ParamVBuild;
			EXEC sp_executesql @SQL, N'@QueryText NVARCHAR(MAX), @ParamList NVARCHAR(MAX)', @QueryText, @ParamList;


			-- Single sp_executesql call with all parameters
            --SET @SQL = N'EXEC sp_executesql N''' + REPLACE(@QueryText, '''', '''''') + ''', N''' + REPLACE(@ParamList, '''', '''''') + ''', ' + @ParamVBuild;
			--PRINT @SQL
            --EXEC(@SQL);

		END

    END
    ELSE
    BEGIN
        -- Execute without parameters
      EXEC sp_executesql @QueryText;
	  PRINT 'no params'
    END
END
