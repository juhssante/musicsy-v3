import { h, render } from 'preact';

import createStore from 'unistore'
import { Provider, connect } from 'unistore/preact'

let store = createStore({ count: 0 })

// If actions is a function, it gets passed the store:
let actions = store => ({
  // Actions can just return a state update:
  increment(state) {
    return { count: state.count + 1 }
  }
})


const App = connect('count', actions)(
  ({ count, increment }) => (
    <div>
      <p>Count: {count}</p>
      <p>{process.env.SERVER_URL}</p>
      <button onClick={increment}>Increment</button>
    </div>
  )
)

const Root = () => (
  <Provider store={store}>
    <App />
  </Provider>
)
render(<Root />, document.getElementById('root'));