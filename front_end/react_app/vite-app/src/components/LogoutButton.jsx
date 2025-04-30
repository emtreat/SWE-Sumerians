import '../App.css'

export function LogoutButton() {
    return (
        <button className="floating-button" style={{
            position: 'fixed',
            color: "white",
            top: '20px',
            left: '20px',
            zIndex: 1000
        }}>
        <i>Logout</i> 
        </button>
    )
}