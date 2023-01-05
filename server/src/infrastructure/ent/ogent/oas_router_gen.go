// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
	}
	if prefix := s.cfg.Prefix; len(prefix) > 0 {
		if strings.HasPrefix(elem, prefix) {
			// Cut prefix from the path.
			elem = strings.TrimPrefix(elem, prefix)
		} else {
			// Prefix doesn't match.
			s.notFound(w, r)
			return
		}
	}
	if len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [1]string{}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'e': // Prefix: "e"
				if l := len("e"); len(elem) >= l && elem[0:l] == "e" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case '-': // Prefix: "-"
					if l := len("-"); len(elem) >= l && elem[0:l] == "-" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 's': // Prefix: "states"
						if l := len("states"); len(elem) >= l && elem[0:l] == "states" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch r.Method {
							case "GET":
								s.handleListEStateRequest([0]string{}, w, r)
							case "POST":
								s.handleCreateEStateRequest([0]string{}, w, r)
							default:
								s.notAllowed(w, r, "GET,POST")
							}

							return
						}
						switch elem[0] {
						case '/': // Prefix: "/"
							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "id"
							// Match until "/"
							idx := strings.IndexByte(elem, '/')
							if idx < 0 {
								idx = len(elem)
							}
							args[0] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								switch r.Method {
								case "DELETE":
									s.handleDeleteEStateRequest([1]string{
										args[0],
									}, w, r)
								case "GET":
									s.handleReadEStateRequest([1]string{
										args[0],
									}, w, r)
								case "PATCH":
									s.handleUpdateEStateRequest([1]string{
										args[0],
									}, w, r)
								default:
									s.notAllowed(w, r, "DELETE,GET,PATCH")
								}

								return
							}
							switch elem[0] {
							case '/': // Prefix: "/event"
								if l := len("/event"); len(elem) >= l && elem[0:l] == "/event" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleReadEStateEventRequest([1]string{
											args[0],
										}, w, r)
									default:
										s.notAllowed(w, r, "GET")
									}

									return
								}
							}
						}
					case 't': // Prefix: "types"
						if l := len("types"); len(elem) >= l && elem[0:l] == "types" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch r.Method {
							case "GET":
								s.handleListETypeRequest([0]string{}, w, r)
							case "POST":
								s.handleCreateETypeRequest([0]string{}, w, r)
							default:
								s.notAllowed(w, r, "GET,POST")
							}

							return
						}
						switch elem[0] {
						case '/': // Prefix: "/"
							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "id"
							// Match until "/"
							idx := strings.IndexByte(elem, '/')
							if idx < 0 {
								idx = len(elem)
							}
							args[0] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								switch r.Method {
								case "DELETE":
									s.handleDeleteETypeRequest([1]string{
										args[0],
									}, w, r)
								case "GET":
									s.handleReadETypeRequest([1]string{
										args[0],
									}, w, r)
								case "PATCH":
									s.handleUpdateETypeRequest([1]string{
										args[0],
									}, w, r)
								default:
									s.notAllowed(w, r, "DELETE,GET,PATCH")
								}

								return
							}
							switch elem[0] {
							case '/': // Prefix: "/event"
								if l := len("/event"); len(elem) >= l && elem[0:l] == "/event" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleReadETypeEventRequest([1]string{
											args[0],
										}, w, r)
									default:
										s.notAllowed(w, r, "GET")
									}

									return
								}
							}
						}
					}
				case 'c': // Prefix: "comments"
					if l := len("comments"); len(elem) >= l && elem[0:l] == "comments" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "GET":
							s.handleListEcommentRequest([0]string{}, w, r)
						case "POST":
							s.handleCreateEcommentRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET,POST")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "id"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							switch r.Method {
							case "DELETE":
								s.handleDeleteEcommentRequest([1]string{
									args[0],
								}, w, r)
							case "GET":
								s.handleReadEcommentRequest([1]string{
									args[0],
								}, w, r)
							case "PATCH":
								s.handleUpdateEcommentRequest([1]string{
									args[0],
								}, w, r)
							default:
								s.notAllowed(w, r, "DELETE,GET,PATCH")
							}

							return
						}
						switch elem[0] {
						case '/': // Prefix: "/"
							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'e': // Prefix: "event"
								if l := len("event"); len(elem) >= l && elem[0:l] == "event" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleReadEcommentEventRequest([1]string{
											args[0],
										}, w, r)
									default:
										s.notAllowed(w, r, "GET")
									}

									return
								}
							case 'u': // Prefix: "user"
								if l := len("user"); len(elem) >= l && elem[0:l] == "user" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleReadEcommentUserRequest([1]string{
											args[0],
										}, w, r)
									default:
										s.notAllowed(w, r, "GET")
									}

									return
								}
							}
						}
					}
				case 'v': // Prefix: "vents"
					if l := len("vents"); len(elem) >= l && elem[0:l] == "vents" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "GET":
							s.handleListEventRequest([0]string{}, w, r)
						case "POST":
							s.handleCreateEventRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET,POST")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "id"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							switch r.Method {
							case "DELETE":
								s.handleDeleteEventRequest([1]string{
									args[0],
								}, w, r)
							case "GET":
								s.handleReadEventRequest([1]string{
									args[0],
								}, w, r)
							case "PATCH":
								s.handleUpdateEventRequest([1]string{
									args[0],
								}, w, r)
							default:
								s.notAllowed(w, r, "DELETE,GET,PATCH")
							}

							return
						}
						switch elem[0] {
						case '/': // Prefix: "/"
							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'a': // Prefix: "admin"
								if l := len("admin"); len(elem) >= l && elem[0:l] == "admin" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleReadEventAdminRequest([1]string{
											args[0],
										}, w, r)
									default:
										s.notAllowed(w, r, "GET")
									}

									return
								}
							case 's': // Prefix: "state"
								if l := len("state"); len(elem) >= l && elem[0:l] == "state" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleReadEventStateRequest([1]string{
											args[0],
										}, w, r)
									default:
										s.notAllowed(w, r, "GET")
									}

									return
								}
							case 't': // Prefix: "type"
								if l := len("type"); len(elem) >= l && elem[0:l] == "type" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleReadEventTypeRequest([1]string{
											args[0],
										}, w, r)
									default:
										s.notAllowed(w, r, "GET")
									}

									return
								}
							case 'u': // Prefix: "users"
								if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleListEventUsersRequest([1]string{
											args[0],
										}, w, r)
									default:
										s.notAllowed(w, r, "GET")
									}

									return
								}
							}
						}
					}
				}
			case 'u': // Prefix: "users"
				if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleListUserRequest([0]string{}, w, r)
					case "POST":
						s.handleCreateUserRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "GET,POST")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch r.Method {
						case "DELETE":
							s.handleDeleteUserRequest([1]string{
								args[0],
							}, w, r)
						case "GET":
							s.handleReadUserRequest([1]string{
								args[0],
							}, w, r)
						case "PATCH":
							s.handleUpdateUserRequest([1]string{
								args[0],
							}, w, r)
						default:
							s.notAllowed(w, r, "DELETE,GET,PATCH")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/events"
						if l := len("/events"); len(elem) >= l && elem[0:l] == "/events" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleListUserEventsRequest([1]string{
									args[0],
								}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					}
				}
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	operationID string
	count       int
	args        [1]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'e': // Prefix: "e"
				if l := len("e"); len(elem) >= l && elem[0:l] == "e" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case '-': // Prefix: "-"
					if l := len("-"); len(elem) >= l && elem[0:l] == "-" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 's': // Prefix: "states"
						if l := len("states"); len(elem) >= l && elem[0:l] == "states" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								r.name = "ListEState"
								r.operationID = "listEState"
								r.args = args
								r.count = 0
								return r, true
							case "POST":
								r.name = "CreateEState"
								r.operationID = "createEState"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
						switch elem[0] {
						case '/': // Prefix: "/"
							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "id"
							// Match until "/"
							idx := strings.IndexByte(elem, '/')
							if idx < 0 {
								idx = len(elem)
							}
							args[0] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								switch method {
								case "DELETE":
									r.name = "DeleteEState"
									r.operationID = "deleteEState"
									r.args = args
									r.count = 1
									return r, true
								case "GET":
									r.name = "ReadEState"
									r.operationID = "readEState"
									r.args = args
									r.count = 1
									return r, true
								case "PATCH":
									r.name = "UpdateEState"
									r.operationID = "updateEState"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}
							switch elem[0] {
							case '/': // Prefix: "/event"
								if l := len("/event"); len(elem) >= l && elem[0:l] == "/event" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch method {
									case "GET":
										// Leaf: ReadEStateEvent
										r.name = "ReadEStateEvent"
										r.operationID = "readEStateEvent"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}
							}
						}
					case 't': // Prefix: "types"
						if l := len("types"); len(elem) >= l && elem[0:l] == "types" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								r.name = "ListEType"
								r.operationID = "listEType"
								r.args = args
								r.count = 0
								return r, true
							case "POST":
								r.name = "CreateEType"
								r.operationID = "createEType"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
						switch elem[0] {
						case '/': // Prefix: "/"
							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "id"
							// Match until "/"
							idx := strings.IndexByte(elem, '/')
							if idx < 0 {
								idx = len(elem)
							}
							args[0] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								switch method {
								case "DELETE":
									r.name = "DeleteEType"
									r.operationID = "deleteEType"
									r.args = args
									r.count = 1
									return r, true
								case "GET":
									r.name = "ReadEType"
									r.operationID = "readEType"
									r.args = args
									r.count = 1
									return r, true
								case "PATCH":
									r.name = "UpdateEType"
									r.operationID = "updateEType"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}
							switch elem[0] {
							case '/': // Prefix: "/event"
								if l := len("/event"); len(elem) >= l && elem[0:l] == "/event" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch method {
									case "GET":
										// Leaf: ReadETypeEvent
										r.name = "ReadETypeEvent"
										r.operationID = "readETypeEvent"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}
							}
						}
					}
				case 'c': // Prefix: "comments"
					if l := len("comments"); len(elem) >= l && elem[0:l] == "comments" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							r.name = "ListEcomment"
							r.operationID = "listEcomment"
							r.args = args
							r.count = 0
							return r, true
						case "POST":
							r.name = "CreateEcomment"
							r.operationID = "createEcomment"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "id"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							switch method {
							case "DELETE":
								r.name = "DeleteEcomment"
								r.operationID = "deleteEcomment"
								r.args = args
								r.count = 1
								return r, true
							case "GET":
								r.name = "ReadEcomment"
								r.operationID = "readEcomment"
								r.args = args
								r.count = 1
								return r, true
							case "PATCH":
								r.name = "UpdateEcomment"
								r.operationID = "updateEcomment"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
						switch elem[0] {
						case '/': // Prefix: "/"
							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'e': // Prefix: "event"
								if l := len("event"); len(elem) >= l && elem[0:l] == "event" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch method {
									case "GET":
										// Leaf: ReadEcommentEvent
										r.name = "ReadEcommentEvent"
										r.operationID = "readEcommentEvent"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}
							case 'u': // Prefix: "user"
								if l := len("user"); len(elem) >= l && elem[0:l] == "user" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch method {
									case "GET":
										// Leaf: ReadEcommentUser
										r.name = "ReadEcommentUser"
										r.operationID = "readEcommentUser"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}
							}
						}
					}
				case 'v': // Prefix: "vents"
					if l := len("vents"); len(elem) >= l && elem[0:l] == "vents" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							r.name = "ListEvent"
							r.operationID = "listEvent"
							r.args = args
							r.count = 0
							return r, true
						case "POST":
							r.name = "CreateEvent"
							r.operationID = "createEvent"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "id"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							switch method {
							case "DELETE":
								r.name = "DeleteEvent"
								r.operationID = "deleteEvent"
								r.args = args
								r.count = 1
								return r, true
							case "GET":
								r.name = "ReadEvent"
								r.operationID = "readEvent"
								r.args = args
								r.count = 1
								return r, true
							case "PATCH":
								r.name = "UpdateEvent"
								r.operationID = "updateEvent"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
						switch elem[0] {
						case '/': // Prefix: "/"
							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'a': // Prefix: "admin"
								if l := len("admin"); len(elem) >= l && elem[0:l] == "admin" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch method {
									case "GET":
										// Leaf: ReadEventAdmin
										r.name = "ReadEventAdmin"
										r.operationID = "readEventAdmin"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}
							case 's': // Prefix: "state"
								if l := len("state"); len(elem) >= l && elem[0:l] == "state" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch method {
									case "GET":
										// Leaf: ReadEventState
										r.name = "ReadEventState"
										r.operationID = "readEventState"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}
							case 't': // Prefix: "type"
								if l := len("type"); len(elem) >= l && elem[0:l] == "type" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch method {
									case "GET":
										// Leaf: ReadEventType
										r.name = "ReadEventType"
										r.operationID = "readEventType"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}
							case 'u': // Prefix: "users"
								if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch method {
									case "GET":
										// Leaf: ListEventUsers
										r.name = "ListEventUsers"
										r.operationID = "listEventUsers"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}
							}
						}
					}
				}
			case 'u': // Prefix: "users"
				if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = "ListUser"
						r.operationID = "listUser"
						r.args = args
						r.count = 0
						return r, true
					case "POST":
						r.name = "CreateUser"
						r.operationID = "createUser"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch method {
						case "DELETE":
							r.name = "DeleteUser"
							r.operationID = "deleteUser"
							r.args = args
							r.count = 1
							return r, true
						case "GET":
							r.name = "ReadUser"
							r.operationID = "readUser"
							r.args = args
							r.count = 1
							return r, true
						case "PATCH":
							r.name = "UpdateUser"
							r.operationID = "updateUser"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/events"
						if l := len("/events"); len(elem) >= l && elem[0:l] == "/events" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: ListUserEvents
								r.name = "ListUserEvents"
								r.operationID = "listUserEvents"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
					}
				}
			}
		}
	}
	return r, false
}
