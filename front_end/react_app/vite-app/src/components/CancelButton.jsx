import '../App.css'

export function CancelButton() {
    return (
        <button className="button" style={{
            position: 'absolute',
            backgroundColor: "#cd6155",
            color: "white",
            top: '20px',
            left: '20px',
            zIndex: 1000
        }}>
        <i className="material-icons">Cancel</i> 
        </button>
    )
}