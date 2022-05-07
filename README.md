# logger
Record logs to the local log file 记录日志到本地的日志文件中


	Instructions: 
  
  Logger is a small tool for logging written in golang. Four types of log statements including info, error, debug, and warning can be written to the specified log file. The user only needs to specify the storage path of the log file (it will be automatically created if it does not exist) and the type and content of the statement to be written to the log. When the size of the log file exceeds 50000, a log_backup folder will be automatically created in the current execution directory. Backup and save the log file to this folder, the backup log file name is the backup time (accurate to microseconds). Then clear the current log file to continue recording new logs.
    
  	使用说明：
  
  logger是用golang编写的一个记录日志的小工具。可以将info、error、debug、warning四种类型的日志语句写入指定的日志文件中。用户只需要指定日志文件的存放路径（若不存在会自动创建）以及要写入日志的语句的类型和内容即可。当日志文件的size超过50000时，会自动在当前执行目录下创建一个log_backup文件夹。将日志文件备份并存入这个文件夹，备份的日志文件名为备份的时间（精确到微秒）。然后清空当前的日志文件，来继续记录新的日志。

	func RecordDebug(filename string, content string)
Write a log statement prefixed with "[DEBUG]" to the specified log file, with the date, time, the filename of the program calling the logger package appended automatically. The parameters to be passed are the log file name and the log statement content.

向指定的日志文件中写入前缀为"[DEBUG]"的日志语句，会自动附加上日期、时间、调用logger的程序文件名。要传入的参数为日志文件名和日志语句内容。

	func RecordWarning(filename string, content string)
	
Write a log statement prefixed with "[WARNING]" to the specified log file, with the date, time, the filename of the program calling the logger package appended automatically. The parameters to be passed are the log file name and the log statement content.

向指定的日志文件中写入前缀为"[WARNING]"的日志语句，会自动附加上日期、时间、调用logger的程序文件名。要传入的参数为日志文件名和日志语句内容。

	func RecordInfo(filename string, content string)
	
Write a log statement prefixed with "[INFO]" to the specified log file, with the date, time, the filename of the program calling the logger package appended automatically. The parameters to be passed are the log file name and the log statement content.

向指定的日志文件中写入前缀为"[INFO]"的日志语句，会自动附加上日期、时间、调用logger的程序文件名。要传入的参数为日志文件名和日志语句内容。

	func RecordError(filename string, content string)
	
Write a log statement prefixed with "[ERROR]" to the specified log file, with the date, time, the filename of the program calling the logger package appended automatically. The parameters to be passed are the log file name and the log statement content.

向指定的日志文件中写入前缀为"[ERROR]"的日志语句，会自动附加上日期、时间、调用logger的程序文件名。要传入的参数为日志文件名和日志语句内容。
