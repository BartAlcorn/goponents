const template = document.createElement("template");
template.innerHTML = `
  <span
      class="todo-item"
      style="display:block"
      _="on load transition #todo-{ID}'s opacity to 1 over 0.4 seconds"
    >
      <span id="btn-done">
        <input
          hx-put="/todos/{ID}/done"
          type="checkbox"
          class="checkbox checkbox-accent checkbox-sm"
        />
      </span>
      <span class="flex-grow"><slot /></span>
      <span
        id="btn-delete"
        class="delete-svg"
        hx-delete="/todos/{ID}"
        hx-target="#todo-{ID}"
        hx-confirm="Are you sure you want to delete this ToDo?"
        _="on htmx:afterRequest transition #todo-{ID}'s opacity to 0 over 0.4 seconds then remove #todo-{ID}"
      >
        <svg xmlns="http://www.w3.org/2000/svg" height="20px" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
          <path stroke-linecap="round" stroke-linejoin="round" d="M12 9.75L14.25 12m0 0l2.25 2.25M14.25 12l2.25-2.25M14.25 12L12 14.25m-2.58 4.92l-6.375-6.375a1.125 1.125 0 010-1.59L9.42 4.83c.211-.211.498-.33.796-.33H19.5a2.25 2.25 0 012.25 2.25v10.5a2.25 2.25 0 01-2.25 2.25h-9.284c-.298 0-.585-.119-.796-.33z" />
        </svg>
      </span>
    </span>
`;

class TodoItem extends HTMLElement {
  constructor() {
    super();

    this.showInfo = true;
    this.ID = this.getAttribute("id");
    this.Checked = this.getAttribute("checked");

    this.attachShadow({ mode: "open" });
    this.shadowRoot.appendChild(template.content.cloneNode(true));
    this.shadowRoot.id = "todo-item-" + this.ID;
    this.shadowRoot
      .querySelector("#btn-done")
      .setAttribute("hx-put", `/todos/${this.ID}/done`);
    this.shadowRoot
      .querySelector("#btn-delete")
      .setAttribute("hx-delete", `/todos/${this.ID}`);
    this.shadowRoot
      .querySelector("#btn-delete")
      .setAttribute("hx-target", `#todo-${this.ID}`);
    this.shadowRoot
      .querySelector("#btn-delete")
      .setAttribute(
        "_",
        `on htmx:afterRequest transition #todo-${this.ID}'s opacity to 0 over 0.4 seconds then remove #todo-${this.ID}`
      );
  }
}

window.customElements.define("todo-item", TodoItem);
