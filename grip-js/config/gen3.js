/*
Reader plugin for reading data from the graph
Example Specimen query:

query{
  Specimen(limit: 100){
    id
    subject {
      reference
    }
    processing {
      method {
        coding {
          code
        }
      }
    }
  }
}
*/
endpoint.add({
    name: "Specimen",
    schema: [{
        id: "String",
        auth_resource_path: "String",
        subject: {
            reference: "String"
        },
        processing: [{
            method: {
                coding: [{
                    code: "String",
                    display: "String",
                    system: "String"
                }],
                text: "String"
            }
        }],
    }],
    args: {
        offset: "Int",
        limit: "Int",
    },
    defaults: {
        offset: 0,
        limit: 100
    },
    handler: (G, args) => {
        return G.V().hasLabel("Specimen").skip(args.offset).limit(args.limit).toList();
    }
})


/*
Mutation plugin for loading customizable schema data via graphql
Example query:

mutation {
  AddSpecimen(
    id: "53c67a06-ea2d-4d24-9249-418dc77a16a9"
    auth_resource_path: "/programs/ohsu/projects/test-invalid"
    collection: {
      bodySite: {
        concept: {
          coding: [{ code: "76752008", display: "Breast", system: "http://snomed.info/sct" }]
          text: "Breast"
        }
      }
      collector: { reference: "Organization/89c8dc4c-2d9c-48c7-8862-241a49a78f14" }
    }
    identifier: [
      { system: "https://my_demo.org/labA", use: "official", value: "specimen_1234_labA" }
    ]
    processing: [
      {
        method: {
          coding: [
            { code: "117032008", display: "Spun specimen (procedure)", system: "http://snomed.info/sct" },
            { code: "Double-Spun", display: "Double-Spun", system: "https://my_demo.org/labA" }
          ]
          text: "Spun specimen (procedure)"
        }
      }
    ]
    resourceType: "Specimen"
    subject: { reference: "Patient/ac4e1aa6-cb52-40e9-8f20-594d9c84f920" }
  ) {
    id
  }
}
*/
endpoint.addMutation({
    name: "AddSpecimen",
    schema: {
        id: "String"
    },
    args: {
        id: "String",
        auth_resource_path: "String",
        collection: {
            bodySite: {
                concept: {
                    coding: [{
                        code: "String",
                        display: "String",
                        system: "String"
                    }],
                    text: "String"
                }
            },
            collector: {
                reference: "String"
            }
        },
        identifier: [{
            system: "String",
            use: "String",
            value: "String"
        }],
        processing: [{
            method: {
                coding: [{
                    code: "String",
                    display: "String",
                    system: "String"
                }],
                text: "String"
            }
        }],
        resourceType: "String",
        subject: {
            reference: "String"
        }
    },
    handler: (G, args) => {
        if (!args.auth_resource_path) {
            throw new Error("auth_resource_path is required");
        }
        const vertexData = {
            auth_resource_path: args.auth_resource_path,
            collection: args.collection || null,
            identifier: args.identifier || null,
            processing: args.processing || null,
            resourceType: args.resourceType || null,
            subject: args.subject || null
        };
        // If you leave args.id blank here it will generate random id for you
        return { "id": G.addVertex(args.id, "Specimen", vertexData) };
    }
});

// another reader plugin
endpoint.add({
    name: "IdsWithSpecimenEdge",
    schema: [{
        id: "String"
    }],
    args: {
        offset: "Int",
        limit: "Int",
    },
    defaults: {
        offset: 0,
        limit: 100
    },
    handler: (G, args) => {
        return G.V().hasLabel("Specimen").out().skip(args.offset).limit(args.limit).toList()
    }
})
