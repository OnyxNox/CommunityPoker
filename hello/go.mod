module example.com/hello

go 1.24.4

replace example.com/greetings => ../greetings

require example.com/greetings v0.0.0-00010101000000-000000000000

require golang.org/x/example/hello v0.0.0-20250605160450-8b405629c4a5
