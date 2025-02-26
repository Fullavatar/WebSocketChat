extends Control


func _on_create_pressed() -> void:
	get_tree().call_group(
		"login",
		"swapUIAuthenticate",
		Global.AuthenticateUIType.SIGNUP,
	)


func _on_forget_pressed() -> void:
	get_tree().call_group(
		"login",
		"swapUIAuthenticate",
		Global.AuthenticateUIType.FORGET,
	)


func _on_login_pressed() -> void:
	pass # Replace with function body.
