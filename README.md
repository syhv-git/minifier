# minifier
This package implements a Minifier for [tdewolff's Minify project](https://github.com/tdewolff/minify). The Minifier takes three parameters, the output dest., the mime type, and the paths to minify. The package will create a new file at the dest. containing the minified program files. 

> The Minifier cannot handle multiple mime types per call, but the Minifier will recursively walk through directories if provided
> 
> The recursive walk will consider all files in the tree to be the same mime type. This can be handled with specific path calls to individual files.
