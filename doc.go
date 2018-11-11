/*

Package detector performs crude network detection by examining the
state of the machine on which it is executing (as opposed to
attempting to access the network). As such the detection should be
both very fast in general and run in approximately the same time
regardless of whether the network is up.

Currently only Linux operating systems are supported.

License

Released under the MIT license.

*/
package detector
