endpoint.add({
    name: "DocumentReference",
    gen3: true,
    schema: [{
        id: "String",
        status: "String",
        resourceType: "String",
        category: [{
            coding: [{
                code: "String",
                display: "String",
                system: "String"
            }]
        }],
        content: [{
            attachment: {
                url: "String"
            }
        }],
        identifier: [{
            system: "String",
            value: "String"
        }]
    }],
    args: {
        offset: "Int",
        limit: "Int",
        project_id : "String"
    },
    defaults: {
        offset: 0,
        limit: 100
    },
    handler: (G, args) => {
        return G.V().hasLabel("DocumentReference").skip(args.offset).limit(args.limit).toList()
    }
})

// Returns all Patient Ids that have a docref edge connected to them.
endpoint.add({
    name: "PatientIdsWithDocumentEdge",
    gen3: true,
    schema: [{
        id: "String"
    }],
    args: {
        offset: "Int",
        limit: "Int",
        project_id : "String"
    },
    defaults: {
        offset: 0,
        limit: 100
    },
    handler: (G, args) => {
        return G.V().hasLabel("DocumentReference").outE("subject").out().skip(args.offset).limit(args.limit).toList()
    }
})

endpoint.add({
    name: "PatientIdsWithSpecimenEdge",
    gen3: true,
    schema: [{
        id: "String"
    }],
    args: {
        offset: "Int",
        limit: "Int",
        project_id : "String"
    },
    defaults: {
        offset: 0,
        limit: 100
    },
    handler: (G, args) => {
        return G.V().hasLabel("Specimen").outE().out().skip(args.offset).limit(args.limit).toList()
    }
})

endpoint.add({
    name: "PatientIdsWithEncounterEdge",
    gen3: true,
    schema: [{
        id: "String"
    }],
    args: {
        offset: "Int",
        limit: "Int",
        project_id : "String"
    },
    defaults: {
        offset: 0,
        limit: 100
    },
    handler: (G, args) => {
        return G.V().hasLabel("Encounter").outE().out().skip(args.offset).limit(args.limit).toList()
    }
})

// {"query":"{  PatientIdsWithDocumentEdge {    id  } }","variables":{}}
// {"query":"query Query($limit: Int){ PatientIdsWithEncounterEdge(limit: $limit) { id }}","variables":{"limit":100000000},"operationName":"Query"}
endpoint.add({
    name: "PatientIdsWithObservationEdge",
    gen3: true,
    schema: [{
        id: "String"
    }],
    args: {
        offset: "Int",
        limit: "Int",
        project_id : "String"
    },
    defaults: {
        offset: 0,
        limit: 100
    },
    handler: (G, args) => {
        return G.V().hasLabel("Observation").outE().out().skip(args.offset).limit(args.limit).toList()
    }
})
