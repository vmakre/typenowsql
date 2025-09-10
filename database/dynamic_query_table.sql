GO

/****** Object:  Table [dbo].[StoredQueries]    Script Date: 10.9.2025 12:46:40 ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [dbo].[TNQ_StoredQueries](
	[QueryID] [int] IDENTITY(1,1) NOT NULL,
	[QueryName] [nvarchar](255) NOT NULL,
	[QueryText] [nvarchar](max) NOT NULL,
	[Description] [nvarchar](max) NULL,
	[LastModified] [datetime] NULL,
	[Parameters] [nvarchar](max) NULL,
	[IsActive] [bit] NULL,
PRIMARY KEY CLUSTERED 
(
	[QueryID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO

ALTER TABLE [dbo].[TNQ_StoredQueries] ADD  DEFAULT (getdate()) FOR [LastModified]
GO

ALTER TABLE [dbo].[TNQ_StoredQueries] ADD  DEFAULT ((1)) FOR [IsActive]
GO

