import express from "express";
import expressGraphQL from "express-graphql";
import mongoose from "mongoose";
import bodyParser from "body-parser";
import cors from "cors";

const app = express();
const PORT = process.env.PORT || "9001";
const db = "mongodb://localhost:27017/inventory";

mongoose
  .connect(db, { useCreateIndex: true, useNewUrlParser: true })
  .then(() => console.log("MongoDB connected"))
  .catch(err => console.error(err));

app.use(
  "/graphql",
  cors(),
  bodyParser.json(),
  expressGraphQL({ schema, graphiql: true })
);

app.listen(PORT, () =>
  console.log(`Server running on http://localhost:${PORT}...`)
);
