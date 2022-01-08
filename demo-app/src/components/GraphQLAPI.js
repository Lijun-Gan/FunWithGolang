import React from "react";
import { ApolloClient, InMemoryCache, ApolloProvider, } from "@apollo/client";
import DisplayData from "./DisplayData";

const GraphQLAPI = () => {

    const client = new ApolloClient({
        cache: new InMemoryCache(),
        uri: "http://localhost:4000/graphql"
    });

    return (
        <div>
            <h1>GraphQLAPI</h1>
            <ApolloProvider client={client}>
                <div>
                    {/* <ul>List of Students</ul> */}
                    <DisplayData />
                </div>
            </ApolloProvider>
        </div>
    )

}

export default GraphQLAPI