package stable

import (
	"fmt"
	"time"
)

func ExampleNew() {
	// create a table
	table := New("table caption")

	// add some field
	table.AddFields("file", "size", "mod")

	// lets add some value
	table.Row("/var/log/docker", 12.453, 0777)

	// print the table
	fmt.Println(table)
	// output:
	// +--------------------------------------+
	// |             table caption            |
	// |--------------------------------------|
	// |        file       |   size   |  mod  |
	// |-------------------+----------+-------|
	// |  /var/log/docker  |  12.453  |  511  |
	// +-------------------+----------+-------+
}

func ExampleSetAlignment() {
	// create a table
	table := New("table caption")

	// add some field
	table.AddFields("file", "size", "mod")

	// change alignment of second and third field
	table.GetField(1).AlignRight()
	// another way
	table.GetField(2).SetAlignment(AlignmentRight)

	// lets add some value
	table.Row("/var/log/docker", 12.453, 0777)

	// print the table
	fmt.Println(table)
	// output:
	// +--------------------------------------+
	// |             table caption            |
	// |--------------------------------------|
	// |        file       |   size   |  mod  |
	// |-------------------+----------+-------|
	// |  /var/log/docker  |  12.453  |  511  |
	// +-------------------+----------+-------+
}

func ExampleSetOption() {
	// create a table
	table := New("table caption")

	// add some field
	table.AddFields("file", "size", "mod")

	// change 'size' fields format option
	table.GetFieldWithName("size").SetOption(&Options{
		Format:    "%0.1f (KB)",
		Alignment: AlignmentCenter,
	})

	table.Row("/var/log/docker", 12.453, 0777)

	fmt.Println(table)
	// output:
	// +-----------------------------------------+
	// |              table caption              |
	// |-----------------------------------------|
	// |        file       |     size    |  mod  |
	// |-------------------+-------------+-------|
	// |  /var/log/docker  |  12.5 (KB)  |  511  |
	// +-------------------+-------------+-------+
}

func ExampleToTable_struct() {
	// example struct
	type Person struct {
		Age    int     `table:"age"`
		Height float64 `table:"height"`
		Name   string  `table:"name"`
		Male   bool    `table:"male"`
	}

	// lets create a person named 'ruby'
	ruby := &Person{
		Name:   "Ruby Cohen",
		Age:    31,
		Height: 1.8,
		Male:   true,
	}

	// convert struct to table
	table, err := ToTable(ruby)
	if err != nil {
		fmt.Println(err)
		return
	}
	// print the table
	fmt.Println(table)
	// output:
	// +------------------------------------------+
	// |                  Person                  |
	// |------------------------------------------|
	// |  age  |  height  |     name     |  male  |
	// |-------+----------+--------------+--------|
	// |  31   |  1.8     |  Ruby Cohen  |  true  |
	// +-------+----------+--------------+--------+
}
func ExampleToTable_anonymous_struct() {
	// lets create an anonymous struct for purpose of example
	// we can change field tag fith 'table' keyword
	fileInfo := struct {
		FilePath string  `table:"path"`
		FileSize float64 `table:"size"`
		FileMod  int     `table:"mod"`
	}{
		FilePath: "/var/log/system.d/docker.log",
		FileSize: 1.8,
		FileMod:  0777,
	}

	// convert struct to table
	table, err := ToTable(fileInfo)
	if err != nil {
		fmt.Println(err)
		return
	}

	// add a caption
	table.SetCaption("docker log file info")

	// lets print 'mod' as octal for more
	table.GetFieldWithName("mod").SetOption(&Options{
		Format: "%o",
	})

	// print the table
	fmt.Println(table)
	// output:
	// +-------------------------------------------------+
	// |               docker log file info              |
	// |-------------------------------------------------|
	// |              path              |  size  |  mod  |
	// |--------------------------------+--------+-------|
	// |  /var/log/system.d/docker.log  |  1.8   |  777  |
	// +--------------------------------+--------+-------+
}
func ExampleToTable_struct_array() {
	// example struct
	type Person struct {
		Age    int     `table:"age"`
		Height float64 `table:"height"`
		Name   string  `table:"name"`
		Male   bool    `table:"male"`
	}

	// lets create a bunch of person
	persons := []*Person{
		{Name: "Ruby Cohen", Age: 30, Height: 1.80, Male: true},
		{Name: "Bethany Parsons", Age: 29, Height: 1.58},
		{Name: "Ronnie Rodriguez", Age: 28, Height: 1.78, Male: true},
		{Name: "Rosa Daniels", Age: 31, Height: 1.80, Male: true},
	}

	// convert struct array to table
	table, err := ToTable(persons)
	if err != nil {
		fmt.Println(err)
		return
	}

	// add a caption
	table.SetCaption("Customers of Coffee Shop")

	// print the table
	fmt.Println(table)
	// output:
	// +-------------------------------------------------+
	// |             Customers of Coffee Shop            |
	// |-------------------------------------------------|
	// |  age  |  height  |        name        |   male  |
	// |-------+----------+--------------------+---------|
	// |  30   |  1.8     |  Ruby Cohen        |  true   |
	// |  29   |  1.58    |  Bethany Parsons   |  false  |
	// |  28   |  1.78    |  Ronnie Rodriguez  |  true   |
	// |  31   |  1.8     |  Rosa Daniels      |  true   |
	// +-------+----------+--------------------+---------+
}
func ExampleToTable_anonymous_struct_array() {
	persons := []*struct {
		Age    int     `table:"age"`
		Height float64 `table:"height"`
		Name   string  `table:"name"`
		Male   bool    `table:"male"`
	}{
		{Name: "Ruby Cohen", Age: 30, Height: 1.80, Male: true},
		{Name: "Bethany Parsons", Age: 29, Height: 1.58},
		{Name: "Ronnie Rodriguez", Age: 28, Height: 1.78, Male: true},
		{Name: "Rosa Daniels", Age: 31, Height: 1.80, Male: true},
	}

	// convert struct array to table
	table, err := ToTable(persons)
	if err != nil {
		fmt.Println(err)
		return
	}

	// add a caption
	table.SetCaption("Customers of Coffee Shop")

	// print the table
	fmt.Println(table)
	// output:
	// +-------------------------------------------------+
	// |             Customers of Coffee Shop            |
	// |-------------------------------------------------|
	// |  age  |  height  |        name        |   male  |
	// |-------+----------+--------------------+---------|
	// |  30   |  1.8     |  Ruby Cohen        |  true   |
	// |  29   |  1.58    |  Bethany Parsons   |  false  |
	// |  28   |  1.78    |  Ronnie Rodriguez  |  true   |
	// |  31   |  1.8     |  Rosa Daniels      |  true   |
	// +-------+----------+--------------------+---------+
}
func ExampleToTable_map() {
	t := time.Date(2022, 01, 17, 0, 0, 0, 0, time.UTC)
	user := map[string]interface{}{
		"username":   "ecoshub",
		"password":   "9b03c12b-ca05-4654-927a-56feb23cb8b3",
		"last_login": t.UnixNano(),
		"region":     "mena",
		"status":     1,
	}

	// convert struct array to table
	table, err := ToTable(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	// add a caption
	table.SetCaption("user info")

	// print the table
	fmt.Println(table)
	// output:
	// +-------------------------------------------------------+
	// |                       user info                       |
	// |-------------------------------------------------------|
	// |      key     |                  value                 |
	// |--------------+----------------------------------------|
	// |  last_login  |  1642377600000000000                   |
	// |  password    |  9b03c12b-ca05-4654-927a-56feb23cb8b3  |
	// |  region      |  mena                                  |
	// |  status      |  1                                     |
	// |  username    |  ecoshub                               |
	// +--------------+----------------------------------------+
}
func ExampleToTable_map_array() {
	t := time.Date(2022, 01, 17, 0, 0, 0, 0, time.UTC)
	user := []map[string]interface{}{
		{
			"username":   "ecoshub",
			"password":   "9b03c12b-ca05-4654-927a-56feb23cb8b3",
			"last_login": t.UnixMilli(),
			"region":     "mena",
			"status":     1,
		},
		{
			"username":   "jenkins99",
			"password":   "981c8036-f017-4b15-920c-4b0c73948cf4",
			"last_login": t.UnixMilli(),
			"region":     "mena",
			"status":     1,
		},
	}

	// convert struct array to table
	table, err := ToTable(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	// add a caption
	table.SetCaption("user info")

	// print the table
	fmt.Println(table)
	// output:
	// +----------------------------------------------------------------------------------------------+
	// |                                           user info                                          |
	// |----------------------------------------------------------------------------------------------|
	// |    last_login   |                password                |  region  |  status  |   username  |
	// |-----------------+----------------------------------------+----------+----------+-------------|
	// |  1642377600000  |  9b03c12b-ca05-4654-927a-56feb23cb8b3  |  mena    |  1       |  ecoshub    |
	// |  1642377600000  |  981c8036-f017-4b15-920c-4b0c73948cf4  |  mena    |  1       |  jenkins99  |
	// +-----------------+----------------------------------------+----------+----------+-------------+
}
func ExampleToTable_json() {
	// example byte array ( json encoded )
	j := []byte(`{"index": 1,"guid": "22c5c5c7-e3b8-40ec-9a83-450bc28c81df","isActive": true,"balance": "$2,057.64","picture": "http://placehold.it/32x32","age": 27}`)

	// convert struct array to table
	table, err := ToTable(j)
	if err != nil {
		fmt.Println(err)
		return
	}

	// add a caption
	table.SetCaption("user info")

	// print the table
	fmt.Println(table)
	// output:
	// +-----------------------------------------------------+
	// |                      user info                      |
	// |-----------------------------------------------------|
	// |     key    |                  value                 |
	// |------------+----------------------------------------|
	// |  age       |  27                                    |
	// |  balance   |  $2,057.64                             |
	// |  guid      |  22c5c5c7-e3b8-40ec-9a83-450bc28c81df  |
	// |  index     |  1                                     |
	// |  isActive  |  true                                  |
	// |  picture   |  http://placehold.it/32x32             |
	// +------------+----------------------------------------+
}
func ExampleToTable_json_array() {
	// example byte array ( json encoded )
	j := []byte(`[{"id": 0,"name": "Heath Vazquez", "age":40, "ssn":"6259d81d221425d39b2b02f5"},{"id": 1,"name": "Blanca Massey", "age":42, "ssn":"6259d8824e829833afacc3c7"},{"id": 2,"name": "Veronica Glass", "age":43, "ssn":"6259d8904d92bf035847e32a"}]`)

	// convert struct array to table
	table, err := ToTable(j)
	if err != nil {
		fmt.Println(err)
		return
	}

	// add a caption
	table.SetCaption("user info")

	// print the table
	fmt.Println(table)
	// output:
	// +--------------------------------------------------------------+
	// |                           user info                          |
	// |--------------------------------------------------------------|
	// |  age  |  id  |       name       |             ssn            |
	// |-------+------+------------------+----------------------------|
	// |  40   |  0   |  Heath Vazquez   |  6259d81d221425d39b2b02f5  |
	// |  42   |  1   |  Blanca Massey   |  6259d8824e829833afacc3c7  |
	// |  43   |  2   |  Veronica Glass  |  6259d8904d92bf035847e32a  |
	// +-------+------+------------------+----------------------------+
}

func ExampleToTable_csv() {
	c := `id,firstname,lastname,email,email2,profession
100,Nikki,Haldas,Nikki.Haldas@yopmail.com,Nikki.Haldas@gmail.com,worker
101,Blinni,Arquit,Blinni.Arquit@yopmail.com,Blinni.Arquit@gmail.com,doctor
102,Shandie,Douglass,Shandie.Douglass@yopmail.com,Shandie.Douglass@gmail.com,firefighter
103,Elie,Phaidra,Elie.Phaidra@yopmail.com,Elie.Phaidra@gmail.com,doctor
104,Jessy,Bahr,Jessy.Bahr@yopmail.com,Jessy.Bahr@gmail.com,police officer
105,Kalina,Hillel,Kalina.Hillel@yopmail.com,Kalina.Hillel@gmail.com,worker
106,Mathilda,Ambrosia,Mathilda.Ambrosia@yopmail.com,Mathilda.Ambrosia@gmail.com,police officer
107,Albertina,Klotz,Albertina.Klotz@yopmail.com,Albertina.Klotz@gmail.com,developer
108,Joeann,Lunsford,Joeann.Lunsford@yopmail.com,Joeann.Lunsford@gmail.com,firefighter
109,Roberta,Moseley,Roberta.Moseley@yopmail.com,Roberta.Moseley@gmail.com,worker
110,Eadie,Riva,Eadie.Riva@yopmail.com,Eadie.Riva@gmail.com,developer
111,Emelina,Keelia,Emelina.Keelia@yopmail.com,Emelina.Keelia@gmail.com,developer
112,Luci,McNully,Luci.McNully@yopmail.com,Luci.McNully@gmail.com,firefighter
113,Aurore,Franza,Aurore.Franza@yopmail.com,Aurore.Franza@gmail.com,doctor
1099,Cissiee,Trey,Cissiee.Trey@yopmail.com,Cissiee.Trey@gmail.com,developer`

	// convert struct array to table
	table, err := CSVToTable(c)
	if err != nil {
		fmt.Println(err)
		return
	}

	// add a caption
	table.SetCaption("user info")

	// print the table
	fmt.Println(table)
	// output:
	// +------------------------------------------------------------------------------------------------------------------------+
	// |                                                        user info                                                       |
	// |------------------------------------------------------------------------------------------------------------------------|
	// |   id   |  firstname  |  lastname  |              email              |             email2            |    profession    |
	// |--------+-------------+------------+---------------------------------+-------------------------------+------------------|
	// |  100   |  Nikki      |  Haldas    |  Nikki.Haldas@yopmail.com       |  Nikki.Haldas@gmail.com       |  worker          |
	// |  101   |  Blinni     |  Arquit    |  Blinni.Arquit@yopmail.com      |  Blinni.Arquit@gmail.com      |  doctor          |
	// |  102   |  Shandie    |  Douglass  |  Shandie.Douglass@yopmail.com   |  Shandie.Douglass@gmail.com   |  firefighter     |
	// |  103   |  Elie       |  Phaidra   |  Elie.Phaidra@yopmail.com       |  Elie.Phaidra@gmail.com       |  doctor          |
	// |  104   |  Jessy      |  Bahr      |  Jessy.Bahr@yopmail.com         |  Jessy.Bahr@gmail.com         |  police officer  |
	// |  105   |  Kalina     |  Hillel    |  Kalina.Hillel@yopmail.com      |  Kalina.Hillel@gmail.com      |  worker          |
	// |  106   |  Mathilda   |  Ambrosia  |  Mathilda.Ambrosia@yopmail.com  |  Mathilda.Ambrosia@gmail.com  |  police officer  |
	// |  107   |  Albertina  |  Klotz     |  Albertina.Klotz@yopmail.com    |  Albertina.Klotz@gmail.com    |  developer       |
	// |  108   |  Joeann     |  Lunsford  |  Joeann.Lunsford@yopmail.com    |  Joeann.Lunsford@gmail.com    |  firefighter     |
	// |  109   |  Roberta    |  Moseley   |  Roberta.Moseley@yopmail.com    |  Roberta.Moseley@gmail.com    |  worker          |
	// |  110   |  Eadie      |  Riva      |  Eadie.Riva@yopmail.com         |  Eadie.Riva@gmail.com         |  developer       |
	// |  111   |  Emelina    |  Keelia    |  Emelina.Keelia@yopmail.com     |  Emelina.Keelia@gmail.com     |  developer       |
	// |  112   |  Luci       |  McNully   |  Luci.McNully@yopmail.com       |  Luci.McNully@gmail.com       |  firefighter     |
	// |  113   |  Aurore     |  Franza    |  Aurore.Franza@yopmail.com      |  Aurore.Franza@gmail.com      |  doctor          |
	// |  1099  |  Cissiee    |  Trey      |  Cissiee.Trey@yopmail.com       |  Cissiee.Trey@gmail.com       |  developer       |
	// +--------+-------------+------------+---------------------------------+-------------------------------+------------------+
}
