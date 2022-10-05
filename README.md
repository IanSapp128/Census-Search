# Census Search
My first attempt are creating a small utility in Go. A utility that pulls up the API documentation for the table you provide it. Supported tables are **Detailed Tables**, **Subject Tables**, and **Data Profiles**.

I initially created this project as a Python script and it had helped me quickly pull up table variables when I needed them. Startup time for the script was pretty slow, so I wanted to try and make it with a compiled language. Should build on Windows, Mac, and Linux but so far only tested on Windows.

**Update**
You can now run this directly by typing csearch *argument*. This will not display any input for the user but just use that passed argument to launch the browser. This is great if you put it in your system's bin file or export a PATH to the directly.

For instance, if you want to look up table DP03 you simply type **csearch DP03** and it should attempt to launch. All rules still apply (Tables that start with B, D, or S). If you type a nonsense table name starting with any of those letters you will just get a 404 on the census page.