import './wasm_exec';
import game from 'url:./game.wasm';

const load = () => {
  if (!WebAssembly.instantiateStreaming) {
    WebAssembly.instantiateStreaming = async (resp, importObject) => {
      const source = await (await resp).arrayBuffer();
      return await WebAssembly.instantiate(source, importObject);
    };
  }
  
  const go = new Go();
  
  WebAssembly.instantiateStreaming(fetch(game), go.importObject).then(result => {
    go.run(result.instance);
  });
}

load();