import { gql } from '@apollo/client';
import { ApolloClient, InMemoryCache } from '@apollo/client';

const client = new ApolloClient({
  uri: 'http://localhost:8890/graphql',
  cache: new InMemoryCache(),
});

const query = async <T>(ql: string): Promise<T> => {
  const { data } = await client.query({
    query: gql(ql),
    fetchPolicy: 'network-only',
  });
  return data;
};

const mutate = async <T>(ql: string, variables: any): Promise<T> => {
  const { data } = await client.mutate({ mutation: gql(ql), variables });
  return data;
};

export { query as baseQuery, mutate as baseMutate };
