TODO:
    20201213: not able to make the go-swagger process the specification that is split into yaml files. need more research.
        found some non go-swagger solutions, but i'm not interested in them.
        with go-swagger I lightly touched the --with-flatten=full option, but could not fugureout how to make it work with multiple yaml files

        the error I'm facing is as below
        $ swagger flatten --with-flatten=full ./swagger/index.yaml
            json: cannot unmarshal string into Go struct field SwaggerProps.definitions of type struct { spec.SchemaProps; spec.SwaggerSchemaProps }


inspired by: 
    https://azimi.me/2015/07/16/split-swagger-into-smaller-files.html
    https://github.com/fireproofsocks/swagger-multifile-examples



read:
    https://www.youtube.com/watch?v=xKhNAQuq1x0

related: 
    https://github.com/go-swagger/go-swagger/issues/1767 about --with-flatten=full option
