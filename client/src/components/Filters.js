import React, { useState } from 'react';

const Filters = ({ onFilterChange }) => {
  const [filterValues, setFilterValues] = useState({
    level: '',
    message: '',
    resourceId: '',
    timestamp: '',
    traceId: '',
    spanId: '',
    commit: '',
    metadata: {
      parentResourceId: '',
    },
  });

  const handleFilterChange = (e) => {
    const { name, value } = e.target;
    if (name.includes('metadata')) {
      // For nested metadata field
      const metadataField = name.split('.')[1];
      setFilterValues({
        ...filterValues,
        metadata: {
          ...filterValues.metadata,
          [metadataField]: value,
        },
      });
    } else {

      setFilterValues({ ...filterValues, [name]: value });
    }
  };

  const applyFilters = () => {
    const filtersToSend = { ...filterValues };
    filtersToSend.metadata = JSON.stringify(filtersToSend.metadata);
    onFilterChange(filtersToSend);
    console.log(filtersToSend)
  };  

  return (
    <div>
      <input
        type="text"
        placeholder="Filter by level..."
        name="level"
        value={filterValues.level}
        onChange={handleFilterChange}
      />
      <input
        type="text"
        placeholder="Filter by message..."
        name="message"
        value={filterValues.message}
        onChange={handleFilterChange}
      />
      <input
        type="text"
        placeholder="Filter by resourceId..."
        name="resourceId"
        value={filterValues.resourceId}
        onChange={handleFilterChange}
      />
      <input
        type="text"
        placeholder="Filter by timestamp(YYYY-MM-DD)..."
        name="timestamp"
        value={filterValues.timestamp}
        onChange={handleFilterChange}
      />
      <input
        type="text"
        placeholder="Filter by traceId..."
        name="traceId"
        value={filterValues.traceId}
        onChange={handleFilterChange}
      />
      <input
        type="text"
        placeholder="Filter by spanId..."
        name="spanId"
        value={filterValues.spanId}
        onChange={handleFilterChange}
      />
      <input
        type="text"
        placeholder="Filter by commit..."
        name="commit"
        value={filterValues.commit}
        onChange={handleFilterChange}
      />
      <input
        type="text"
        placeholder="Filter by parentResourceId..."
        name="metadata.parentResourceId" // Adjust the name for nested field
        value={filterValues.metadata.parentResourceId} // Update value for nested field
        onChange={handleFilterChange}
      />
      <button onClick={applyFilters}>Filter Search</button>
    </div>
  );
};

export default Filters;
