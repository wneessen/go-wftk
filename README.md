# Winni's Go Web-Framework Toolkit
This package provides a list of useful Go functions, that I regualarly use in my web applications. 
The functions are seperated into different sub packages and can be imported on their own

* **crypto**: Everything around cryptography
  * **argon2**: Helper functions to handle Argon2id hashing
  * **random**: Functions that provide cryptographically secure generator functions
* **error**: All about error handling
  * **json**: Provides convenient functions to provide standarized JSON errors
* **geoip**: Helper functions for Maxmind's GeoIP database (Requires valid database)
* **password**: Helper functions to make passwords more secure (Strength check, HIBP)
  