package repo

//go:generate sh -c "rm -rf mocks && mkdir mocks"
//go:generate minimock -i Repo -o ./mocks/ -s "_minimock.go"
