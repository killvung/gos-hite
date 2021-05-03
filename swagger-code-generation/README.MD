# Strava V3 Swagger Code Generation Example

To generate Strava V3 library in Go lang. Follow the [Strava API V3 Documentation](https://developers.strava.com/docs/#client-code) to get the package installed, then run the command `swagger-codegen geneate --input-spec https://developers.strava.com/swagger/swagger.json --lang go --output generated/` to generated boilerpalte code and stored in `generated/`

## Notes

For macOS user, after running `brew install swagger-codegen@2 maven`, make sure to identify the location of the package through `brew list swagger-codegen@2`. For some reasons, brew doesn't set alias for the package located at `/usr/local/Cellar/swagger-codegen@2/2.4.19/bin/swagger-codegen` by default, strange.
