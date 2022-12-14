import Foundation

func startLoad() {
    let sem = DispatchSemaphore.init(value: 0)

    let url = URL(string: "https://www.example.com/")!
    let task = URLSession.shared.dataTask(with: url) { data, response, error in
        defer { sem.signal() }

        if let error = error {
            print("\(error)")
            return
        }
        guard let httpResponse = response as? HTTPURLResponse,
            (200...299).contains(httpResponse.statusCode) else {
            print("\(response!)")
            return
        }
        if let mimeType = httpResponse.mimeType, mimeType == "text/html",
            let data = data,
            let string = String(data: data, encoding: .utf8) {

            print(string)
        }
    }
    task.resume()

    sem.wait()
}

startLoad()
