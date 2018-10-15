class Table extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            images: []
        }
    }

    componentDidMount() {
        fetch('/goserver')
            .then((response) => {
                return response.json()
            })
            .then((data) => {
                this.setState({images: data})
            });
    }

    render() {
        let rows = this.state.images.map((row) => {
            return (<Row item={row} key={row.id}/>)
        });

        return (
            <div className="container">
                <h1>Data Entry</h1>
                <table className="table table-striped">
                    <thead>
                    <tr>
                        <th>Id</th>
                        <th>Title</th>
                        <th>Media</th>
                        <th>Artist</th>
                        <th>Gallery</th>
                        <th>Organization</th>
                    </tr>
                    </thead>
                    <tbody>
                    {rows}
                    </tbody>
                </table>
            </div>
        )
    }
}
