import Surreal from "surrealdb.js"
import Express from "express"


const db = new Surreal("http://localhost:8000/rpc")
const app = Express()
const main = async () => {

	try {
		// Signin as a namespace, database, or root user
		await db.signin({
			user: 'root',
			pass: 'root',
		});
		await db.use("test", "test");
		const info = await db.info();



		app.get("/", async (req, res) => {
			let people = await db.select("todo");
			res.json(people);
		})
		app.listen("3000", () => console.log("Connected to Muala"));
	} catch (err) {

		console.error(err)
	}
}

main()
const createTodo = async (/** @type {string} */ name, /** @type {Boolean} */ done) => {
	let created = await db.create("todo", {
		name, done
	});
}
const deleteTodo = async (/** @type {string} */ name) => {
	await db.delete("todo", {
		name
	})
}