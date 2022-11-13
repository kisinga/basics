const {parse} = require ("@babel/parser")
const traverse = require ("@babel/traverse").default

const code = "2 + (4 * 10)";

const ast = parse(code)

traverse(ast, {
    // NumericLiteral(path) {
    //     console.log(path.node.value)
    // }
    NumericLiteral: {
        enter(path) {
            console.log(`Entered ${path.node.value}`)
        },
        exit(path) {
            console.log(`Exited ${path.node.value}`)
        }
    }
})