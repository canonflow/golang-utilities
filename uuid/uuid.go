package uuid

import "uuid"

func UUID() string {
	return uuid.NewV7().String()
}

func UUIDV4() string {
	return uuid.NewV4().String()
}

func UUIDV7() string {
	return uuid.NewV7().String()
}

func UUIDMin() string {
	return uuid.Nil().String()
}

func UUIDMax() string {
	return uuid.Max().String()
}
