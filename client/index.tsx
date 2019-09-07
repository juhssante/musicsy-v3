import { h, render } from 'preact'
import { Provider } from 'unistore/preact'
import Router from './router'

import createStore from 'unistore'

const store = createStore({ count: 0 })

const app = document.getElementById('root')

render(
  <Provider store={store}>
    <Router />
  </Provider>,
  app
)