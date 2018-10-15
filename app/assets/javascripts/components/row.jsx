class Row extends React.Component {

    constructor(props) {
        super(props);
    }

    render() {
        let {item} = this.props;

        return (
            <tr>
                <td>{item.id}</td>
                <td>{item.title}</td>
                <td>{item.media} ({item.width}x{item.height})</td>
                <td>{item.artist}</td>
                <td>{item.gallery}</td>
                <td>{item.organization}</td>
            </tr>
        )
    }
}
