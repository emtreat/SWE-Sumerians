import '../App.css'

export function AddFile() {
    return (
        <button className="floating-button" style={{
          position: 'fixed',
          color: "white",
          bottom: '20px',
          right: '20px',
          zIndex: 1000
      }}>
        <i className="material-icons">Add File</i> 
      </button>
    )
}