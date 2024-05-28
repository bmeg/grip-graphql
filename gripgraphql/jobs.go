 package gripgraphql

 import(

    "fmt"
    "reflect"

    "github.com/bmeg/grip/gripql"
 )

func graphStmtsEqual(stmts1 []*gripql.GraphStatement, stmts2 []*gripql.GraphStatement) bool {
    if len(stmts1) != len(stmts2){
        return false
    }
    for i := range stmts1 {
        fmt.Println("REFLECT DEEP EQUAL: ", reflect.DeepEqual(stmts1[i], stmts2[i]), stmts1[i], stmts2[i])
        if !(reflect.TypeOf(stmts1[i].GetStatement()) == reflect.TypeOf(stmts2[i].GetStatement()) &&  reflect.DeepEqual(stmts1[i], stmts2[i])){
            return false
        }
    }
    return true
}

func (cw *JSClientWrapper) GetCachedJob(query gripql.GraphQuery, CachedGS []*gripql.GraphStatement, RemainingGS []*gripql.GraphStatement) (chan *gripql.QueryResult, error) {
    status, err := cw.client.ListJobs(cw.graph)
    if err != nil{
        fmt.Println("ERR: ", err)
        return nil, err
    }
 
    var result chan *gripql.QueryResult;
    var jobList []string;
    for _, elem := range status{
        jobList = append(jobList, elem.Id)
    }
    fmt.Println("JOB LIST: ", jobList)
 
    jobHit := false
    for _, jobId := range jobList{
        job, err := cw.client.GetJob(cw.graph, jobId)
        if err != nil{
            fmt.Printf("ERR: %s", err)
            return nil, err
        }
        fmt.Printf("JOB: %#v\n",job)
        if graphStmtsEqual(job.Query, CachedGS){
            jobHit = true
            fmt.Println("RESUMING JOB")
            // check to make sure that the job is finished before resuming
            if job.State == 2{
                result, err = cw.client.ResumeJob(job.Graph, job.Id, &gripql.GraphQuery{Graph: job.Graph, Query: RemainingGS})
                if err != nil{
                    fmt.Println("ERR: ", err)
                    return nil, err
                }
                fmt.Println("OUT RESUMED JOB FINISHED")
                return result, nil
            }
        }
    }
    if jobHit == false{
         submit, err := cw.client.Submit(&gripql.GraphQuery{Graph: cw.graph, Query: CachedGS })
         if err != nil{
             fmt.Printf("ERR: %s", err)
             return nil, err
         }
         fmt.Printf("Job not found submitting new job: %#v\n", submit)
    }
    return nil, nil
}


