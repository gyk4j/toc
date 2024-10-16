import {
    MDBModal,
    MDBModalDialog,
    MDBModalContent,
    MDBModalHeader,
    MDBModalTitle,
    MDBModalBody,
    MDBModalFooter,
    MDBBtn
} from 'mdb-react-ui-kit';

export default function PreTocCheckerDialog({ 
    showPreTocCheckerDialog,
    closePreTocChecker,
    fileStatus,
    webStatus,
    mailStatus,
    databaseStatus,
    actionText
    }) {

    const COLOR = {
        "Pending":      { fg: "text-black", bg: "bg-light" },
        "Running":      { fg: "text-white", bg: "bg-info" },
        "Completed":    { fg: "text-white", bg: "bg-success" },
        "Failed":       { fg: "text-white", bg: "bg-danger" },
    }

    return (
        <MDBModal staticBackdrop show={showPreTocCheckerDialog} tabIndex='-1'>
            <MDBModalDialog size="lg">
                <MDBModalContent>
                    <MDBModalHeader>
                        <MDBModalTitle>Pre-TOC Checker</MDBModalTitle>
                        <MDBBtn className='btn-close' color='none' onClick={closePreTocChecker}></MDBBtn>
                    </MDBModalHeader>
                    <MDBModalBody>
                        <div className='container-fluid'>
                            <div className='row'>
                                <div className='col-md-10'>File</div>
                                <div className={`col-md-2 ${ COLOR[fileStatus].fg } bg-opacity-75 ${ COLOR[fileStatus].bg }`}>
                                    {fileStatus}
                                </div>
                            </div>
                            <div className='row'>
                                <div className='col-md-10'>Web</div>
                                <div className={`col-md-2 ${ COLOR[webStatus].fg } bg-opacity-75 ${ COLOR[webStatus].bg }`}>
                                    {webStatus}
                                </div>
                            </div>
                            <div className='row'>
                                <div className='col-md-10'>Mail</div>
                                <div className={`col-md-2 ${ COLOR[mailStatus].fg } bg-opacity-75 ${ COLOR[mailStatus].bg }`}>
                                    {mailStatus}
                                </div>
                            </div>
                            <div className='row'>
                                <div className='col-md-10'>Database</div>
                                <div className={`col-md-2 ${ COLOR[databaseStatus].fg } bg-opacity-75 ${ COLOR[databaseStatus].bg }`}>
                                    {databaseStatus}
                                </div>
                            </div>
                        </div>
                    </MDBModalBody>
                    <MDBModalFooter>
                        <MDBBtn onClick={closePreTocChecker} color="secondary">Exit</MDBBtn>
                        <MDBBtn onClick={closePreTocChecker}>{actionText}</MDBBtn>
                    </MDBModalFooter>
                </MDBModalContent>
            </MDBModalDialog>
        </MDBModal>
    )
}