

# exec stored procedure

```
USE [sakila]
GO
DECLARE	@return_value int
EXEC	@return_value = [dbo].[sp_TNQ_ExecuteStoredQueryWithJSON]
		@QueryName = N'query4',
		@ParamInput = N'{"id":1}'

SELECT	'Return Value' = @return_value
GO
```