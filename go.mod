module tdd

go 1.15

replace test.com/dog => ./animals/dog

replace test.com/cat => ./animals/cat

replace test.com/human => ./animals/human

replace test.com/supertype => ./supertype

require (
	test.com/cat v0.0.0-00010101000000-000000000000
	test.com/dog v0.0.0-00010101000000-000000000000
	test.com/human v0.0.0-00010101000000-000000000000
	test.com/supertype v0.0.0-00010101000000-000000000000
)
