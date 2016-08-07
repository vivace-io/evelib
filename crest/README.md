# Package crest

A CREST client library implementation, complete with rate limiting.

## Usage

Start using the package, you only need to execute the package's Init() function.

 - `userAgent` is sent with every request to CREST
 - `root` is the base CREST URI. The package provides TranquilityURI and SingularityURI as suggestions, but you can easily pass your own URI (a proxy, for example). However, the proxy should follow the same endpoint layout as CREST All URI's should have a trailing slash.
 - `rate` is the maximum request rate per second. Must be within the domain [1,150].
 - `burst` is the maximum request burst rate. Must be within the domain [150,400].

After that, calls to retrieve models and data from CREST may be executed.


## Return Types

Each method will generally return a pointer or a slice of pointers to various CREST models. The main reason behind this is, depending on the model in question, that it is made possible for the model to call Complete() and Walk() on its self to update its information locally.
