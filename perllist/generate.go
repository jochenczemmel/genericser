package perllist

// code generation
// float64
//go:generate genericser -in perllist.go -out perllist_float64.go -r "T1=float64,T2=MyFloat64"
// Person
//go:generate genericser -in perllist.go -out perllist_person.go -r "T1=Person,T2=PersonList"
