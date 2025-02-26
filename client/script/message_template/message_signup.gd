extends Message
class_name SignUp

@export var username:String
@export var email:String
@export var password:String

func _init() -> void:
	type = Global.MessageType.SIGNUP


func to_json() -> String:
	var message_dict = {
		"Type": type_to_string(),
		"Username": username,
		"Email": email,
		"Password": password,
	}
	return JSON.stringify(message_dict, "\t")
