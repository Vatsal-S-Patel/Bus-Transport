export function CustomizeTable({ details, handleDelete, items, isRoute }) {

    console.log(isRoute)
  return (
    <div className="mt-8">
      <hr></hr>
      <table className="min-w-full divide-y divide-gray-200">
        <thead className="bg-gray-50">
          <tr>
            {Object.keys(details).map((data, index) => {
              if (isRoute && index === Object.keys(details).length - 1) {
                return ("")
              }
              return (
                <th
                  key={index}
                  scope="col"
                  className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  {data}
                </th>
              );
            })}
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Actions
            </th>
          </tr>
        </thead>
        <tbody className="bg-white divide-y divide-gray-200">
          {items != undefined
            ? items.map((item, index) => (
                <tr key={index}>
                  {Object.keys(item).map((data) => {
                    return (
                      <td className="px-6 py-4 whitespace-nowrap">
                        {item[data]}
                      </td>
                    );
                  })}
                  <td>
                    <button
                      onDoubleClick={() => handleDelete(item.id, index)}
                      className="text-red-500 hover:text-red-700"
                    >
                      Delete
                    </button>
                  </td>
                </tr>
              ))
            : ""}
        </tbody>
      </table>
    </div>
  );
}
