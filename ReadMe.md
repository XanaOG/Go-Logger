# Go-Logger

## Usage

### Add Colors
**Custom Color Addition:**  
If you wish to use custom colors, please ensure to add them in an earlier function before attempting to call the display function with the custom color.

To add colors, you can utilize `GoLogger.AddColor(Tag (string), ColorCode (string))`.  
**Example:**  
`GoLogger.AddColor("CUSTOM", "\033[1;95m")`  
This example results in a bright purple color being added with the tag "CUSTOM".

### Log
To invoke the logging function, utilize `GoLogger.Display(Tag (string), Content (string))`.

**Examples:**  
- `GoLogger.Display("INFO", "Test")`  
  This would result in an Info log.  
- `GoLogger.Display("CUSTOM", "Test")`  
  This would result in a Test message with the "Custom" tag and color.

For color example please refer to the [Colors File](https://github.com/XanaOG/GoLogger/blob/main/Colors.md) (Courtesy of [iamnewton](https://github.com/iamnewton/) for the basic examples). The app supports any ANSI colors however, so you can use whatever colors you like.