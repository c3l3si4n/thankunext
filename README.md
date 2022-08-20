# thankunext
Easily gather all routes related to a NextJs application through parsing of _buildManifest.js

This is useful because NextJS will only load JS files that are needed in that page, making it so some funcionality becomes hidden from devTools. Extracting all endpoints defined in buildManifest.js allows you to:
- Quickly grab all routes defined in the application
- Get all .js chunks that are used in the app, and then download then later for examination.

# Installation
```
go install github.com/c3l3si4n/thankunext@latest
```

# How does it work?

1. Receive input from argument, consider that URL to be pointing to an application made in NextJS.
2. Check if given URL ends in "buildManifest.js", if it does then evaluate the buildManifest.js script in a headless browser to construct the self.__BUILD_MANIFEST variable.
3. If buildManifest.js is already loaded (such as in the root of the application) or if it was evalued in the last step, extract the contents of self.__BUILD_MANIFEST variable to golang's runtime.
4. Output NextJS config contents to STDOUT


