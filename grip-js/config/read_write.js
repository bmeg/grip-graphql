endpoint.add({
    name: "Prompts",
    //cached: true,
    schema: [{
        id: "String",
        prompt: "String"
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
        return G.V().hasLabel('Prompt').skip(args.offset).limit(args.limit).render({"id":"$._id", "prompt":"$.prompt"}).toList()
    }
})

endpoint.addMutation({
    name: "AddPrompt",
    schema: {
        id: "String"
    },
    args: {
        prompt: "String"
    },
    handler: (G, args) => {
        return { "id" : G.addVertex(null, "Prompt", {"prompt" : args.prompt} ) }
    }
})
