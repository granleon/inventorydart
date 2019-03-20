export default `
type User {
  id: String
  username: String!
}

type Query {
  user(id: String!): User
  users: [User]
}

type Mutation {
  addUser(id: String, username: String!): User
  editUser(id: String, username: String): User
  deleteUser(id: String, username: String): User
}
`;
