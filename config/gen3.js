endpoint.add({
    name: "PatientIdsWithDocumentEdge",
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
        return G.V().hasLabel("Encounter").outE().out().skip(args.offset).limit(args.limit).toList()
    }
})

// {"query":"{  PatientIdsWithDocumentEdge {    id  } }","variables":{}}
// {"query":"query Query($limit: Int){ PatientIdsWithEncounterEdge(limit: $limit) { id }}","variables":{"limit":100000000},"operationName":"Query"}
endpoint.add({
    name: "PatientIdsWithObservationEdge",
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
