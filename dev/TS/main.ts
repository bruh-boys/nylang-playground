await (async function main() {
  let output: string = '🍙 main = 🏨 () {\n     🎤🎶 ( "hello world" ) ;\n } ;\n';

  if (window.location.pathname != "/")
    output = document.getElementById("real-code")!.innerText;

  // @ts-ignore
  var editor = CodeMirror(
    document.getElementById("notebook-content"),
    {
      lineNumbers: true,
      mode: "javascript",
      value: output,
      theme: "material",
    }
  );

  const codeMirro: Element = document.getElementsByClassName("CodeMirror")[0]
  codeMirro.id = "code-editor";

  document.getElementById("share")!.addEventListener("click", async (): Promise<void> => {
    // @ts-ignore
    const code = editor.getValue();

    const res: Response = await fetch("/share", {
      body: code,
      method: "POST",
    });

    window.location.href = await res.text();
  });

  var reader: any;

  async function sendPost(e: any) {
    try {
      await reader.cancel()

    } catch (_) { };

    document.getElementById('output-code')!.innerText = "";

    const resp: Response = await fetch(`/${e.srcElement.id}`, {
      method: "POST",
      // @ts-ignore
      body: await editor.getValue(),
    });

    reader = resp.body!.getReader();

    while (true) {
      const { done, value } = await reader.read();

      if (done) break;

      let out: string = String.fromCharCode.apply(null, value)

      document.getElementById('output-code')!.innerText += out;
    }
  }

  document.getElementById("stop")!.addEventListener("click", async (): Promise<void> => {
    if (reader) try {
      await reader.cancel();

    } catch (_) { };

  });

  document.getElementById("run")!.addEventListener("click", sendPost);
  document.getElementById("lexer")!.addEventListener("click", sendPost);
  document.getElementById("parser")!.addEventListener("click", sendPost);

})();

export {};