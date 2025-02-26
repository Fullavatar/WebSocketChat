extends RichTextLabel
class_name Logger

func write(logs:Global.LogType):
	append_text(type_to_string(logs) + "\n")


func type_to_string(type:Global.LogType) -> String:
	match type:
		Global.LogType.NOT_VALID_ADDRESS:
			return "Please enter a valid address/port"
		Global.LogType.CONNECTING:
			return "Connecting to server..."
		Global.LogType.CONNECTION_FAILED:
			return "Failed to connect to server"
		Global.LogType.CONNECTION_TIMEOUT:
			return "Connection timeout"
		Global.LogType.CANNOT_PARSE_RESPONSE:
			return "Can't parse response"
		Global.LogType.CONNECTION_ESTABLISHED:
			return "Server reached"
		Global.LogType.CONNECTION_SUCCESS:
			return "Connection to server is successful"
	return "Unknown"
