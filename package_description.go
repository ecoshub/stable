/*

Stable can create ascii table from;

	-   structs                     (struct)
	-   struct arrays               ([]string)
	-   json encoded byte arrays    ([]byte)
	-   string interface maps       (map[string]interface{})
	-   string interface map arrays ([]map[string]interface{})
	-   csv encoded strings         (string)
	-   custom row by row values    (.Row(values...interface{}))


Functionalities:

	-   wide range of type support
	-   value and header orientation options
	-   custom print format option
	-	char limiting
	-   customizable border styles
	-	and much more...

*/
package stable
