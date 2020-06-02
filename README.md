# reactima-slate-go
Package for server side html rendering of internal json format from incredible [Slatejs](https://github.com/ianstormtaylor/slate) editor 

### Editor Example 
https://www.slatejs.org/

```javascript
const SlateEditor = props => {
  const { input = {} } = props
  const { values = {} } = input

  const noteText = values.note
    ? deserialize(values.note)
    : [
      {
        children: [{ text: "" }]
      }
    ]
  const initialValue = values.noteSlate ? values.noteSlate : noteText

  const [value, setValue] = useState(initialValue)

  useEffect(() => {
    if (props.input && props.input.values && !props.input.values.noteSlate) {
      setValue([
        {
          children: [{ text: "" }]
        }
      ])
    }
  }, [props])

  const renderElement = useCallback(props => <Element {...props} />, [])
  const renderLeaf = useCallback(props => <Leaf {...props} />, [])
  const editor = useMemo(() => withImages(props.id, props.type)(withHtml(withLinks(withHistory(withReact(createEditor()))))), [])

  const { debug = false } = props

  if (!props.input || !props.input.onChange || !props.input.name) {
    return (
      <div className="form-group row">
      <div className="alert alert-danger w-100">
      Broken input form "{props.name}
    ":
    {(!props.input || !props.input.onChange) && " no input.onChange"}
    {(!props.input || !props.input.name) && " no input.name"}
    {!props.type && " no type"}
  </div>
    </div>
  )
  }

  return (
    <div className="mr-0 ml-0 mt-1 mb-1">
    <Error>
    <Slate
  editor={editor}
  value={value}
  onChange={value => {
    setValue(value)

    const { values, name, onChange = () => {} } = props.input

    let inputData = {}
    inputData[name + "Slate"] = value
    inputData[name] = extractText(value)

    onChange({}, inputData)
  }}
>
<div className={""}
  style={{
    position: "relative",
      padding: "1px 0px 0px 1px",
      margin: "0px 0px 0px 0px",
      borderBottom: "1px solid #eee",
      marginBottom: 10
  }}
>
<MarkButtonReactima format="bold" icon="hh-icon-edit-bold hh-icon-sm " />
    <MarkButtonReactima format="italic" icon="hh-icon-edit-italic hh-icon-sm" />
    <MarkButtonReactima format="underlined" icon="hh-icon-edit-underline hh-icon-sm" />
    <MarkButtonReactima format="code" icon="hh-icon-edit-code hh-icon-sm" />
    <BlockButtonReactima format="heading-one" icon="hh-icon-edit-h1 hh-icon-sm" />
    <BlockButtonReactima format="heading-two" icon="hh-icon-edit-h2 hh-icon-sm" />
    <BlockButtonReactima format="block-quote" icon="hh-icon-edit-paragraph hh-icon-sm" />
    <BlockButtonReactima format="numbered-list" icon="hh-icon-edit-list-ol hh-icon-sm" />
    <BlockButtonReactima format="bulleted-list" icon="hh-icon-edit-list-ul hh-icon-sm" />
    <LinkButton />
    <InsertImageButton />
    </div>
    <Editable
  renderElement={renderElement}
  renderLeaf={renderLeaf}
  placeholder="Enter some rich text…"
  spellCheck
  autoFocus
  style={{ backgroundColor: "white", minHeight: 200 }}
  onKeyDown={event => {
    for (const hotkey in HOTKEYS) {
      if (isHotkey(hotkey, event)) {
        event.preventDefault()
        const mark = HOTKEYS[hotkey]
        toggleMark(editor, mark)
      }
    }
  }}
  />
  <ReactimaJsonPrint data={{ value }} />
  </Slate>
  </Error>
  </div>
)
}

```

### JSON OutPut
* See [test file](slate_test.go)