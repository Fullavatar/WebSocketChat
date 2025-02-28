extends GutTest

func test_pass():
	assert_eq(2, 1 + 1)


func test_fail():
	assert_false(true)
