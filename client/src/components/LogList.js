import React from 'react';

const LogList = ({ logs }) => {
  if (!logs || !Array.isArray(logs)) {
    return <div>No logs available.</div>;
  }

  return (
    <div>
      <h2>Logs</h2>
      <div style={{ display: 'flex', flexWrap: 'wrap' }}>
        {logs.map((log, index) => (
          <div key={index} style={{ border: '1px solid #ccc', borderRadius: '5px', padding: '20px', margin: '30px', width: '300px' }}>
            <p><strong>Level:</strong> {log.level}</p>
            <p><strong>Message:</strong> {log.message}</p>
            <p><strong>Resource ID:</strong> {log.resourceId}</p>
            <p><strong>Timestamp:</strong> {log.timestamp}</p>
            <p><strong>Trace ID:</strong> {log.traceId}</p>
            <p><strong>Span ID:</strong> {log.spanId}</p>
            <p><strong>Commit:</strong> {log.commit}</p>
            <p><strong>Metadata:</strong></p>
            <ul style={{ listStyleType: 'none' }}>
              {Object.keys(log.metadata).map((key, i) => (
                <li key={i}>
                  <strong>{key}:</strong> {log.metadata[key]}
                </li>
              ))}
            </ul>
          </div>
        ))}
      </div>
    </div>
  );
};

export default LogList;
