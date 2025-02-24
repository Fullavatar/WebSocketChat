extends RichTextLabel
class_name Logger

func write(logs:Global.LogType):
	append_text(type_to_string(logs) + "\n")


func type_to_string(type:Global.LogType) -> String:
	match type:
		Global.LogType.VALIDADDRESS:
			return "Please enter a valid address/port"
		Global.LogType.CONNECTING:
			return "Connecting to server..."
		Global.LogType.FAILEDCONNECTION:
			return "Failed to connect to server"
		Global.LogType.TIMEOUT:
			return "Connection timeout"
		Global.LogType.CANNOTPARSE:
			return "Can't parse response"
		Global.LogType.CONNECTIONSUCCESS:
			return "Connection established"
	return "Unknown"
