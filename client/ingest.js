    const numberOfLogs = 1000; // Change this to the desired number of log ingestion requests
    const logs = [];

    // Create 10 log data objects with different timestamps
    for (let i = 0; i < 10; i++) {
        const hour = (i+1).toString().padStart(2, '0'); // Varying hours
        const timestamp = `2023-09-1${i}T${hour}:00:00Z`; // Timestamp format
        const logData = {
            "level": "info",
            "message": `Log ${i + 1}`,
            "resourceId": `server-${i + 1}`,
            "timestamp": timestamp,
            "traceId": `trace-${i + 1}`,
            "spanId": `span-${i + 1}`,
            "commit": `commit-${i + 1}`,
            "metadata": {
                "parentResourceId": `parent-server-${i + 1}`
            }
        };
        logs.push(logData);
    }

    async function ingestLogs() {
        const startTime = Date.now();
        for (let i = 0; i < numberOfLogs; i++) {
            const logToIngest = logs[i % 10]; // Use logs in a cycle for ingestion
            try {
                const fetch = await import('node-fetch');
                const response = await fetch.default('http://localhost:3000', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(logToIngest),
                });

                if (!response.ok) {
                    console.error(`Request ${i + 1} failed with status: ${response.status}`);
                }
            } catch (error) {
                console.error(`Error making request ${i + 1}:`, error);
            }
        }
        const endTime = Date.now();
        const totalTime = endTime - startTime;
        console.log(`Ingested ${numberOfLogs} logs in ${totalTime}ms`);
    }

    ingestLogs();
