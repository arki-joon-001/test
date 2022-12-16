// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import Arkijoon001TestTest from './arkijoon001.test.test'


export default { 
  Arkijoon001TestTest: load(Arkijoon001TestTest, 'arkijoon001.test.test'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}