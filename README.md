# minifier
This package implements a Minifier for [tdewolff's Minify project](https://github.com/tdewolff/minify). The Minifier takes three parameters, the output dest., the mime type, and the paths to minify. The package will create a new file at the dest. containing the minified program files. 

The Minifier performs a recursive walk on directories provided in the variadac string parameter. All files are read into memory and after all paths have been walked, the contents are minified and written to the output file.

> The minifier will ignore the destination file if it exists in the tree being enumerated
> 
> The Minifier will not throw an error when handling multiple mime types per call (but your projects will ;)