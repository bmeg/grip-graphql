endpoint.add({
    name: "PatientIdsWithDocumentEdge",
    //cached: true,
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
        return G.V().hasLabel('DocumentReference').out('subject').skip(args.offset).limit(args.limit).toList()
    }
})

endpoint.add({
    name: "PatientIdsWithSpecimenEdge",
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
        return G.V().hasLabel("Specimen").out().skip(args.offset).limit(args.limit).toList()
    }
})

endpoint.add({
    name: "PatientIdsWithEncounterEdge",
    //cached: false,
    schema: [{
        id: "String"
    }],
    args: {
        offset: "Int",
        limit: "Int",
        project_id : "String",

    },
    defaults: {
        offset: 0,
        limit: 100
    },
    handler: (G, args) => {
        return G.V().hasLabel("Encounter").out('subject_Patient').skip(args.offset).limit(args.limit).toList()
    }
})

// {"query":"{  PatientIdsWithDocumentEdge {    id  } }","variables":{}}
// {"query":"query Query($limit: Int){ PatientIdsWithEncounterEdge(limit: $limit) { id }}","variables":{"limit":100000000},"operationName":"Query"}
endpoint.add({
    name: "PatientIdsWithObservationEdge",
    //cached: false,
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
        return G.V().hasLabel("Observation").out().skip(args.offset).limit(args.limit).toList()
    }
})
