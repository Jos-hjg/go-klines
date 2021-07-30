function GetRecordList(){
    let newlyId = 0
    let datas = []
    let url = "http://localhost:8082/records"
    $.ajax(url, {
        type: "GET",
        success: (data) => {
            data.data.forEach(function (element) {
                let dd = []
                newlyId = element["ID"]
                //console.log("newlyId:", newlyId)
                dd.push(element["MicroTime"])
                dd.push(element["Open"])
                dd.push(element["Highest"])
                dd.push(element["Lowest"])
                dd.push(element["Close"])
                dd.push(element["Volume"])
                datas.push(dd)
            })

        }
    })
    return datas
}

function  GetCurrentRecord(datas){
    let dd = []
    let url = "http://localhost:8082/record"
    $.ajax(url,{
        type: "GET",
        success: (data) => {
            //console.log("compare: ", data.data["ID"], newlyId)
            if(data.data["ID"] !== newlyId && newlyId !== 0) {
                newlyId = data.data["ID"]
                dd.push(data.data["MicroTime"])
                dd.push(data.data["Open"])
                dd.push(data.data["Highest"])
                dd.push(data.data["Lowest"])
                dd.push(data.data["Close"])
                dd.push(data.data["Volume"])
                datas.push(dd)
            }else if(data.data["ID"] === newlyId){
                datas.pop()
                        
                dd.push(data.data["MicroTime"])
                dd.push(data.data["Open"])
                dd.push(data.data["Highest"])
                dd.push(data.data["Lowest"])
                dd.push(data.data["Close"])
                dd.push(data.data["Volume"])
                datas.push(dd)
            }

        }
        
    })
    return datas
}

