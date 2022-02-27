## Bitmap

A simple and easy-to-use bitmap library based on Go language.

## Usage

```go
m1 := Bitmap{}
m1.Add(1)
m1.Add(10)
m1.Add(999)

m2 := Bitmap{}
m2.Add(1)
m2.Add(10)
m2.Add(998)

m1.IntersectWith(&m2)

fmt.Printf("%v", m1.Elems())
```

## API

- ```Add``` ：add the element to the set
- ```Remove``` ：remove the element in the set
- ```Clear``` ：make the bitmap empty
- ```Copy``` ：copy the bitmap to a new bitmap
- ```Has``` ：whether the element already exists in the set
- ```Elems``` ：return the every element in the set
- ```UnionWith``` ：calculate the unionSet
- ```IntersectWith``` ：calculate the interSet
- ```DifferenceWith``` ：calculate the differenceSet

## License
MIT
