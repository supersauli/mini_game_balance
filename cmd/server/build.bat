go build -gcflags="all=-N -l" -ldflags="-compressdwarf=false" -o angela-member-center.exe
set goos=linux
go build -gcflags="all=-N -l" -ldflags="-compressdwarf=false" -o angela-member-center