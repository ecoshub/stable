# Welcome to Simple Table (stable)

## What is 'stable'?
stable can create ascii table from;
-   structs
-   struct arrays
-   json encoded byte arrays
-   string interface maps
-   string interface map arrays
-   csv encoded strings
-   custom row by row values

## Functionalities
-   wide range of type support
-   value and header orientation options
-   custom print format option
-	char limiting
-   customizable border styles
-	and much more...

## Try your self!
anonymous struct array example:
```go
	persons := []struct {
		Name   string
		Age    int
		Height float64
		Male   bool
	}{
		{Name: "Ruby Cohen", Age: 30, Height: 1.80, Male: true},
		{Name: "Bethany Parsons", Age: 29, Height: 1.58, Male:false},
		{Name: "Ronnie Rodriguez", Age: 28, Height: 1.78, Male: true},
		{Name: "Rosa Daniels", Age: 31, Height: 1.80, Male: true},
	}

    // convert it to table
	t, err := stable.ToTable(persons)
	if err != nil {
		fmt.Println(err)
		return
	}

    // set the table caption
	t.SetCaption("Customers")

    // print the table
	fmt.Println(t)

// output: 
// +-------------------------------------------------+
// |                    Customers                    |
// |-------------------------------------------------|
// |        Name        |  Age  |  Height  |   Male  |
// |--------------------+-------+----------+---------|
// |  Ruby Cohen        |  30   |  1.8     |  true   |
// |  Bethany Parsons   |  29   |  1.58    |  false  |
// |  Ronnie Rodriguez  |  28   |  1.78    |  true   |
// |  Rosa Daniels      |  31   |  1.8     |  true   |
// +--------------------+-------+----------+---------+
```

## Create custom tables!
```go
	// create a table with caption
	table := stable.New("Benchmark of Hashing")

	// add fields
	table.AddFields(
		"File",
		"File (KB)",
		"Chunk (KB)",
		"Time (ms)",
	)

	// add a field with more option
	table.AddFieldWithOptions("Prop (%)", &stable.Options{
		Format:         "%0.2f",
		Alignment: stable.AlignmentCenter,
	})

	// add row
	table.Row("/var/log/system/crontab.log", 12.515, 14.265, "32", 0.223)
	table.Row("/var/log/system/monit.log", 85.521, 43.32, nil, 0.742)
	table.Row("/var/log/system/logrotate.log", 96.57, nil, "31112", 0.321)
	table.Row("/var/log/system/docker-daemon.log", 13.3511, 34.01, "3652", 0.895)

	// print the table
	fmt.Println(table)
	// output:
	// +---------------------------------------------------------------------------------------------+
	// |                                     Benchmark of Hashing                                    |
	// |---------------------------------------------------------------------------------------------|
	// |                 File                |  File (KB)  |  Chunk (KB)  |  Time (ms)  |  Prop (%)  |
	// |-------------------------------------+-------------+--------------+-------------+------------|
	// |  /var/log/system/crontab.log        |  12.515     |  14.265      |  32         |    0.22    |
	// |  /var/log/system/monit.log          |  85.521     |  43.32       |  -          |    0.74    |
	// |  /var/log/system/logrotate.log      |  96.57      |  -           |  31112      |    0.32    |
	// |  /var/log/system/docker-daemon.log  |  13.3511    |  34.01       |  3652       |    0.90    |
	// +-------------------------------------+-------------+--------------+-------------+------------+

```
## Whats Next?
-	custom border style
