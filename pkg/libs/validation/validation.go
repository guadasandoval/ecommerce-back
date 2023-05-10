package validation

// Rule type
type Rule uint

// Enum of all rules
const (
	InvalidRule = iota

	List           // Array that need to validate each item
	Struct         // Struct that need to validate
	Required       // Value can not be NULL or empty
	Same           // Value must match with other field of the list - Param0: Name of the othe field
	Alpha          // Check only alphabetic characters
	Numeric        // Check only numeric characters
	AlphaNumeric   // Check only alphabetic and numeric characters
	Int            // Check if the string is an integer
	IntNotRequired // Check if the string is an integer

	MinValue // Valid if number if greater or equal to Param0 - Param0: number to compare
	MaxValue // Valid if number if lower or equal to Param0 - Param0: number to compare
	// Valid if number if greater or equal to Param0 and lower or equal to Param1
	// Param0: number to compare, Param1: number to compare
	BetweenValue

	Length    // Check if the string length is equal to Param0 - Param0: length
	MinLength // Check if the string length is greater or equal to Param0 - Param0: min length
	MaxLength // Check if the string length is lower or equal to Param0 - Param0: max length

	Date     // Expect a valid Date
	Time     // Expect a valid Time
	DateTime // Expect a valid Date Time

	EMail // Valid EMail

	In // Valid if in the list - Param0: value1, Param1: value2 ...
	// Check if exists in database or if Param2 exists, is equal to Param2
	// Param0: table name, Param1: column name, Param2: valid value
	DBExists
	DBNotExists // Check if not exists in database - Param0: table name, Param1: column name
	// Param0: table name, Param1: column name, Param2: exclude column name, Param3: exclude field name to get the value
	DBUnique

	OptionParam                // Check if the option exists on the database
	OptionParamNotRequired     // Check if the option exists on the database
	OptionParamList            // Check if the option exists on the database
	OptionParamListNotRequired // Check if the option exists on the database
)

var (
	mapRules map[string]Rule
	mapNames map[Rule]string
)

func init() {
	// const op errors.Operation = "pkg.libs.validation.init"

	mapRules = make(map[string]Rule)
	mapNames = make(map[Rule]string)

	mapRules["list"] = List
	mapNames[List] = "List"
	mapRules["struct"] = Struct
	mapNames[Struct] = "Struct"
	mapRules["required"] = Required
	mapNames[Required] = "Required"
	mapRules["same"] = Same
	mapNames[Same] = "Same"
	mapRules["alpha"] = Alpha
	mapNames[Alpha] = "Alpha"
	mapRules["numeric"] = Numeric
	mapNames[Numeric] = "Numeric"
	mapRules["alphanumeric"] = AlphaNumeric
	mapNames[AlphaNumeric] = "AlphaNumeric"
	mapRules["int"] = Int
	mapNames[Int] = "Int"
	mapRules["intnotrequired"] = IntNotRequired
	mapNames[IntNotRequired] = "IntNotRequired"

	mapRules["minvalue"] = MinValue
	mapNames[MinValue] = "MinValue"
	mapRules["maxvalue"] = MaxValue
	mapNames[MaxValue] = "MaxValue"
	mapRules["betweenvalue"] = BetweenValue
	mapNames[BetweenValue] = "BetweenValue"

	mapRules["length"] = Length
	mapNames[Length] = "Length"
	mapRules["minlength"] = MinLength
	mapNames[MinLength] = "MinLength"
	mapRules["maxlength"] = MaxLength
	mapNames[MaxLength] = "MaxLength"

	mapRules["date"] = Date
	mapNames[Date] = "Date"
	mapRules["time"] = Time
	mapNames[Time] = "Time"
	mapRules["datetime"] = DateTime
	mapNames[DateTime] = "DateTime"

	mapRules["email"] = EMail
	mapNames[EMail] = "EMail"

	mapRules["in"] = In
	mapNames[In] = "In"
	mapRules["dbexists"] = DBExists
	mapNames[DBExists] = "DBExists"
	mapRules["dbnotexists"] = DBNotExists
	mapNames[DBNotExists] = "DBNotExists"
	mapRules["dbunique"] = DBUnique
	mapNames[DBUnique] = "DBUnique"

	mapRules["optionparam"] = OptionParam
	mapNames[OptionParam] = "OptionParam"
	mapRules["optionparamnotrequired"] = OptionParamNotRequired
	mapNames[OptionParamNotRequired] = "OptionParamNotRequired"
	mapRules["optionparamlist"] = OptionParamList
	mapNames[OptionParamList] = "OptionParamList"
	mapRules["optionparamlistnotrequired"] = OptionParamListNotRequired
	mapNames[OptionParamListNotRequired] = "OptionParamListNotRequired"
}
